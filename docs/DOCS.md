# Webfetch docs 

## Configuration

### Frontend

**webfetch uses the user's frontend**.
it has basic frontend on the start, but the users can write their own frontends.

**how to write my own frontend?** 
you need to go to the your config directory (e.g. ~/.config) and just write what you want.  
in the *index.html* you can find a lines like a {{.Platform}} - yeah, you're right,
you should add these lines in the your frontend.

**note: you can't use typescript, only javascript** 

### Backend

today you can change only one thing in the backend, it's a port.
to change port you need to go to the your config directory,
find a *config.toml* file, open it and change the port number **with a :**  

"4242" - **WRONG!!!!!!**  
":4242" - **yeah sweety :3** 
