package sessions

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"storiesservice/internal/consts"
	"storiesservice/pkg/signer"
	"storiesservice/pkg/utctime"
	"storiesservice/pkg/varsizedint"
	"time"
	"unsafe"
)

func New(userId int64) string {
	unixNano := uint64(utctime.UnixNano())
	sessionRaw := make([]uint8, unsafe.Sizeof(unixNano)+unsafe.Sizeof(userId))
	binary.LittleEndian.PutUint64(sessionRaw, unixNano)
	sessionRaw = sessionRaw[:int(unsafe.Sizeof(unixNano))+varsizedint.Encode(sessionRaw[unsafe.Sizeof(unixNano):], uint64(userId))]
	return hex.EncodeToString(append(signer.Sign(sessionRaw, consts.UserSessionPepper[:]), sessionRaw...))
}

func Parse(session string) (userId int64, err error) {
	var sessionRaw []uint8
	sessionRaw, err = hex.DecodeString(session)
	if err != nil {
		err = errors.New(consts.I18nSessionDecodeFailed)
		return
	}

	if !signer.Verify(sessionRaw[signer.HashSize:], consts.UserSessionPepper[:], sessionRaw[:signer.HashSize]) {
		err = errors.New(consts.I18nSessionVerifyFailed)
		return
	}

	unixNano := int64(binary.LittleEndian.Uint64(sessionRaw[signer.HashSize:]))
	if utctime.Get().After(time.Unix(0, unixNano).UTC().Add(consts.UserSessionTime)) {
		err = errors.New(consts.I18nSessionExpired)
		return
	}

	userId = int64(varsizedint.Decode(sessionRaw[signer.HashSize+unsafe.Sizeof(uint64(0)):]))
	return
}
