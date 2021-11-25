## <% .ProjectName %>

### Application commands
* `start-server` - start a web application:
  ```bash
  ./application start-server
  ```
* `migrate` - migrate changes to the database:
  ```bash
  ./application migrate
  ```

For more information about application commands, run:
```bash
# list all available commands
./application

# get command usage info
./application [command] --help
```

## Build and run the docker container.
```bash
docker build -t <% .ProjectName | lower %>:latest .
docker run -p 8000:8000 <% .ProjectName | lower %>:latest ./application start-server --bind 0.0.0.0:8000 --workers=5
```
