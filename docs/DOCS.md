# Webfetch docs 

## Configuration

### Frontend

**webfetch uses your custom frontend**.
it comes with a basic frontend by default, but you can write your own one.

**how do i write my own frontend?**   
go to the your config directory (e.g., ~/.config/webfetch) and create the frontend files
you want.  
in *index.html* you'll find placeholders like a {{.Platform}}.
you're right, you can use these in your custom frontend to display system info.

**note: you can't use typescript, only javascript is supported** 

### Backend

currently, you can change one thing in the backend: the port.  
to change it, go to the your config directory,
open a *config.toml* file and set the port **with a colon**:

"4242" - **!!!!!!WRONG!!!!!!**  
":4242" - **yeah sweety, good boy :3** 
