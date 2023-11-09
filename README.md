
# ASCII - ART - WEB - EXPORT - FILE
* `ayessenb` 
* `roshakba` 



Ascii-art-web-export-file consists on making your site :

    more appealing, interactive and intuitive.
    more user friendly.
    give more feedback to the user.
    has Dockerfile for creating multiple images and containers
    additional feature which also allows to export the file with the resulted ascii-art


#### Description:


* We used text/template pakage to recieve and send GET/POST requests.
* GET /: Sends HTML response, the main page.
* POST /ascii-art: sends data to Go server (text and a banner). Displays the result in the route /ascii-art after the POST is completed. So going from the home page to another page.
* Linked external css file to make good design 
* Added a Dockerfile which creates two images and container



## Usage/Examples
Cloning storage to your host
```CMD/Terminal 
git clone git@github.com:compygirl/ascii-art-web-export-file.git
```
Go to the downloaded repository:

```CMD/Terminal 
cd ascii-art-web-export-file
```
Run a program:
```CMD/Terminal 
go run .
```

Follow the link on the terminal:
```CMD/Terminal 
Starting server got testing... http://127.0.0.1:8080 
```


Choose your font and write a message to output



## HTTP status codes
* OK (200), if everything went without errors.
* Not Found, if nothing is found, for example templates or banners.
* Bad Request, for incorrect requests.
* Internal Server Error, for unhandled errors.



## Feedback

If you liked our project, we would be grateful if you could add `Star` to the repository.

Alem Student
25.07.2023.