## <% .ProjectName %>

### Application commands
* `start-server` - start a web application on the local machine:
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

### Build and run
Build the Docker container and run the application:
```bash
sudo docker build -t <% .ProjectName | to_snake_case %>:latest .
docker run -p 8000:8000 <% .ProjectName | to_snake_case %>:latest ./application start-server --bind 0.0.0.0:8000 --workers=5
```
