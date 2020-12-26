# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.define "docker" do | d |
     d.vm.box = "gusztavvargadr/docker-linux"
     d.vm.network "forwarded_port", guest: 80, host: 8080
     d.vm.network "forwarded_port", guest: 80, host: 8080
     d.vm.network "private_network", ip: "192.168.10.10"
     # d.vm.network "public_network"  
     d.vm.synced_folder ".", "/home/vagrant/app"
     d.vm.provider "virtualbox" do |vb|
   # Display the VirtualBox GUI when booting the machine
       vb.name = "Docker"
       vb.gui = false
       # Customize the amount of memory on the VM:
       vb.memory = "2048"
     end
     d.vm.provision "shell", inline: <<-SHELL
        apt-get update
     SHELL
  end
end