# StoriesService
The project is an Api for third-party integrations, such as stories of different formats, both photo and video
## Features
- Authentication
- Authorisation
- Create, view, edit, delete stories
- Select the type of story to publish (image/video)
- Editing content before publishing
- Create multiple stories at once
  
## Installation

```bash
git clone https://github.com/f0rxz/StoriesService
```

# Process

In internal/consts/consts_nrelease.go and internal/consts/consts_release.go change the characteristics:
- port
- user
- dbname

# Start
``` bash
cd StoriesService-main
./build.sh    
./server.out
```
