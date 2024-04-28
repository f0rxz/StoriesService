package main

import (
	"storiesservice/internal/db"
	"text/template"
)

var tmplBaseBegin = template.Must(template.New("base_begin").Parse(`<!DOCTYPE html>
<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=no">
		<style>
			* {
				font-family: verdana;
				font-size: 16px;
			}

			body {
				background-color: #101010;
				color: #f0f0f0;
			}

			a {
				border-radius: 16px;
				color: #7070f0;
				text-decoration: none;
			}

			fieldset {
				border-radius: 16px;
			}

			input {
				min-height: 24px;
			}
		</style>
	</head>
	<body bgcolor="#101010">
		<font color="f0f0f0" face="verdana" size="3">
			<center>
				<a href="?lang=ru"><font color="#7070f0">Русский</font></a> <a href="?lang=en"><font color="#7070f0">English</font></a><br>
`))

var tmplBaseEnd = template.Must(tmplBaseBegin.New("base_end").Parse(`
			</center>
		</font>
	</body>
</html>
`))

var tmplErrMsg = template.Must(tmplBaseEnd.New("sign_in").Parse(`
	{{template "base_begin"}}
	<br><font color="red">{{.}}</font>
	<meta http-equiv="refresh" content="3; url=/">
	{{template "base_end"}}
`))

var tmplMsg = template.Must(tmplBaseEnd.New("sign_in").Parse(`
	{{template "base_begin"}}
	<br><font color="lime">{{.}}</font>
	<meta http-equiv="refresh" content="3; url=/">
	{{template "base_end"}}
`))

type signInData struct {
	Legend        string
	LabelUsername string
	LabelPassword string
	Submit        string
}

var tmplSignIn = template.Must(tmplBaseEnd.New("sign_in").Parse(`
{{template "base_begin"}}
<br><table width="300" border="0" cellspacing="0" cellpadding="5">
	<tr>
		<td>
			<form action="/sign_in" method="post">
				<fieldset>
					<legend>{{.Legend}}</legend>
					<label for="username">{{.LabelUsername}}:</label><br>
					<input type="text" id="username" name="username" required><br><br>
					<label for="password">{{.LabelPassword}}:</label><br>
					<input type="password" id="password" name="password" required><br><br>
					<input type="submit" value="{{.Submit}}">
				</fieldset>
			</form>
		</td>
	</tr>
</table>
{{template "base_end"}}
`))

var tmplStoryStyles = template.Must(tmplBaseEnd.New("base_story").Parse(`
<style>
	.story-container {
		width: 270px;
		height: 480px;
		background-color: #202020;
		margin-bottom: 10px;
		overflow: hidden;
		display: flex;
		justify-content: center;
		align-items: center;
		cursor: pointer;
		user-select: none;
		position: relative;
		text-align: left;
	}

	.story-container img, .story-container video {
		max-width: 100%;
		max-height: 100%;
		width: auto;
		height: auto;
	}

	.story-container span {
		position: absolute;
		font-weight: bold;
		word-break: break-word;
		text-shadow: 0px 0px 10px black;
		-webkit-text-stroke: 1px black;
	}

	#addStory {
		width: 50px; 
		height: 50px;
		font-size: 24px; 
		line-height: 50px; 
		text-align: center; 
		background-color: lightblue; 
		border: none; 
		cursor: pointer; 

		display:block;
	}
</style>
<script>
	function storyTextResize(textPosX, textPosY, storyText) {
		storyText.style.fontSize = Math.min(32, Math.floor(Math.sqrt(270 * 480 / storyText.textContent.length / 2))).toString() + 'px';
		storyText.style.left = Math.floor(textPosX - storyText.clientWidth / 2).toString() + 'px';
		storyText.style.top = Math.floor(textPosY - storyText.clientHeight / 2).toString() + 'px';
	}
</script>
`))

type storyModifyData struct {
	EndpointValue string
	TypeValue     int16
	LinkValue     string
	TextValue     string
	TextPosXValue int
	TextPosYValue int
	Legend        string
	LabelType     string
	OptionImage   string
	OptionVideo   string
	LabelLink     string
	LabelText     string
	LabelTextPos  string
	Submit        string
}

var tmplStoryModify = template.Must(tmplStoryStyles.New("story_modify").Parse(`
{{template "base_begin"}}
{{template "base_story"}}
<br><table width="300" border="0" cellspacing="0" cellpadding="5">
	<tr>
		<td>
			<form action="{{.EndpointValue}}" method="post">
				<fieldset>
					<legend>{{.Legend}}</legend>
					<label for="type">{{.LabelType}}:</label><br>
					<select id="type" name="type">
						{{if eq .TypeValue 0}}
							<option value="image" selected>{{.OptionImage}}</option>
							<option value="video">{{.OptionVideo}}</option>
						{{else}}
							<option value="image">{{.OptionImage}}</option>
							<option value="video" selected>{{.OptionVideo}}</option>
						{{end}}
					</select><br>
					<label for="link">{{.LabelLink}}:</label><br>
					<input type="text" id="link" name="link" value="{{.LinkValue}}" required><br>
					<label for="text">{{.LabelText}}:</label><br>
					<input type="text" id="text" name="text" value="{{.TextValue}}" required maxlength="128"><br>
					<label>{{.LabelTextPos}}:</label><br>
					<input type="hidden" id="textPosXInput" name="text_pos_x" value="{{.TextPosXValue}}">
					<input type="hidden" id="textPosYInput" name="text_pos_y" value="{{.TextPosYValue}}">
					
					<div id="story" class="story-container">
						<span id="storyText"></span>
						<script>
							var x = parseInt(textPosXInput.value), y = parseInt(textPosYInput.value);
							storyText.textContent = text.value;
							storyTextResize(x, y, storyText);
							var dragStartX, dragStartY;
							var dragging = false;
							storyText.addEventListener('mousedown', function(evt) {
								dragStartX = evt.clientX - parseInt(textPosXInput.value);
								dragStartY = evt.clientY - parseInt(textPosYInput.value);
								dragging = true;
							});
							window.addEventListener('mouseup', function(evt) {
								dragging = false;
							});
							story.addEventListener('mousemove', function(evt) {
								if (evt.buttons === 0) {
									dragging = false;
								}
								if (dragging) {
									var x = evt.clientX - dragStartX, y = evt.clientY - dragStartY;
									textPosXInput.value = Math.floor(x).toString();
									textPosYInput.value = Math.floor(y).toString();
									storyTextResize(x, y, storyText);
								}
							});
							for (var eventType of ['change', 'keydown', 'keypress', 'keyup']) {
								text.addEventListener(eventType, function(evt) {
									var x = Math.floor(270 / 2), y = Math.floor(480 / 2);
									textPosXInput.value = x.toString();
									textPosYInput.value = y.toString();
									storyText.textContent = text.value;
									storyTextResize(x, y, storyText);
								});
							}
						</script>
					</div><br>
					<input type="submit" value="{{.Submit}}">
				</fieldset>
				
			</form>

		</td>
		
	</tr>
</table>
<button id="addStory">+</button>
<script>
	b=document.querySelector("#addStory");
	b.addEventListener("click", ()=>{
		document.querySelector("tr").innerHTML += document.querySelector("td").innerHTML
	});
</script>
				
{{template "base_end"}}
`))

type indexData struct {
	IsAdmin      bool
	AStoryAdd    string
	AStoryEdit   string
	AStoryRemove string
	Stories      []*db.Story
}

var tmplIndex = template.Must(tmplStoryStyles.New("index").Parse(`
{{template "base_begin"}}
{{template "base_story"}}
{{if .IsAdmin}}
<br><a href="/story_add"><font color="#7070f0">{{.AStoryAdd}}</font></a><br>
{{end}}
<br>
{{range .Stories}}
	{{if $.IsAdmin}}
		<a href="/story_edit?id={{.Id}}"><font color="#7070f0">{{$.AStoryEdit}}</font></a>
		<details style="display: inline;">
			<summary>+</summary>
			<a href="/story_remove?id={{.Id}}"><font color="#7070f0">{{$.AStoryRemove}}</font></a>
		</details>
		<br>
	{{end}}
	{{if eq .Type 0}}
	<div class="story-container">
		<img src="{{.Link}}">
	{{else if eq .Type 1}}
	<div class="story-container" onclick="this.children[0].play();">
		<video src="{{.Link}}"></video>
	{{end}}
		<span id="storyText{{.Id}}">{{.Text}}</span>
		<script>
			storyTextResize({{.TextPosX}}, {{.TextPosY}}, storyText{{.Id}});
		</script>
	</div>
{{end}}
{{template "base_end"}}
`))
