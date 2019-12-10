cd update_ssh_config
echo "Compile update_ssh_config [arm5]"
env GOOS=linux GOARCH=arm GOARM=5 go build -o ~/update_ssh_config_linux_arm
echo "Compile update_ssh_config [amd64]"
env GOOS=linux GOARCH=amd64 go build -o ~/update_ssh_config_linux_amd64