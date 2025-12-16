# TXP

## Configuration

### TXP Project Structure

    my_project/  <- - - Root project folder, is mounted to /txp_data when running in a container
        txpfile.yml  <- - - Project configuration
        fonts/
            (your font files here)...
        templates/
            demo/  <- - - Note: templates must have their own subfolders
                main.typ


### curl example call

    	`curl -X POST 'https://exampleserver.com/template/{template_name}' \
    	--header 'Content-Type: application/json' \
    	--data '{ "message":"Hello, World!" }'`
