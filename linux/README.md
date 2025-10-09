<h3>Linux <img width="50" height="50" alt="image" src="https://github.com/user-attachments/assets/97810118-3a39-490b-802f-c1c77208ff67" />
</h3> 

This repo contains helpful notes and useful linux commands that I’ve found useful during my learning journey.

### Server
<p>
  serve something like,

1. file server serve file
2. web server serves static pages
3. database server serve data

server serve info and client request for that info

[youtube.com](http://youtube.com) - this is domain

it has application server - serve dynamic data

standalone apps - don’t need internet

web app - need application

operating system - helps running programs

majority applications run on linux

unix - macOS (paid)

linux is free, open source
</p>

<hr>

### Linux and windows
- window - commercial license
- linux - GPL - general public license
- linux - dve, programming, scripting
- window - general purpose
- linux - no antivirus need

<hr>

### Linux system architecture
Kernel - contains everything to talk to hardware - cpu, ram

Shell - platform to communicate with kernel using commands

Boot loader - starts system and runs kernel

<hr>

### Linux commands
Everything in linux is treated as file

It all starts from `/` —> root folder

- `cd` command - “change directory”

```bash
cd /
```

- `ls` - list all items

```bash
ls

>> Applications	etc		private		Users
bin		home		sbin		usr
cores		Library		System		var
dev		opt		tmp		Volumes
```

- `date`

```bash
date

>> Mon Oct  6 15:34:42 IST 2025
```

- `mkdir` - make directory

```bash
cd documents/workspace
mkdir devops
```

- creating file

```bash
cd devops
touch newFile.txt
```

- `pwd` - present working directory

```bash
pwd

>> /Users/neelkumar/documents/workspace/devops
```

- going back

```bash
pwd
>> /Users/neelkumar/documents/workspace/devops
cd ..
pwd
>> /Users/neelkumar/documents/workspace
```

- `rm` - remove file

```bash
cd devops
rm newFile.txt
```

- `rmdir` - remove folder

```bash
mkdir mydir
rmdir mydir
```

- adding content in file

```bash
touch intro.txt
echo "hello world" > intro.txt
```

- getting content of file

```bash
cat intro.txt

>> hello world
```

- getting 1st 10 lines

```bash
head intro.txt
```

- getting last 10 line

```bash
tail intro.txt
```

- copying a file into folder

⇒`cp source target`

```bash
cp intro.txt myFolder/

cp workspace/devops/files/newFile.txt workspace/devops/myCloud/
```

- to copy all files of one folder to other folder

`-r` means recursively

```bash
cp -r folder1/ folder2/
```

- moving files

```bash
mv style.css myCloud/
```

- renaming a folder

`devops` to `linux`

```bash
cd ..
mv devops/ linux
```

- word count

```bash
wc intro.txt
>>        2       6      28 intro.txt
```

lines = 2

words = 6

bytes = 28

- Hard link
    
    A **hard link** is like having *two names for the same file*.
    
    - Both the original file and the hard link **point to the same data** on disk.
    - If you delete one, the other still works — because the actual data is still there.
    - Changes in one are seen in the other, since they’re the same file internally.
- Soft link
    
    A **soft link** is like a *shortcut* or *pointer* to another file’s path.
    
    - It stores the **path** of the target file, not its actual data.
    - If the original file is deleted, the symlink becomes **broken**

- creating shortlink

```bash
ln -s email.txt shortlink.txt   
```

- creating hard link

```bash
ln password.txt secrets.txt
```

- `cut`

content inside bio.txt:

> “my name is neel”
> 

```bash
cut -b 1-2 bio.txt
>> my
```

- `tee` - add content in file and also prints it

```bash
echo "xyz@gmail.com" | tee email.txt
```

- sorting content in data

```bash
sort items.txt
```

- difference b/w 2 files

```bash
diff email.txt password.txt

>> 1c1
< xyz@gmail.com
---
> password: 123
```

- disk file storage check

```bash
df -h
```
<hr>

### User and group management

### System level commands

- `uname` - tells platform we are using

```bash
uname
>> Linux
```

- The `uptime` command in Linux provides information about how long a system has been running since its last boot

```bash
uptime
>>  7:11  up  1:01, 2 users, load averages: 1.30 1.46 1.56
```

- `who` gives info of all users and their login activity

```bash
who
>> 
neelkumar        console       7 Oct 06:11 
neelkumar        ttys000       7 Oct 06:56 
```

- `whoami` - tells about current user

```bash
whoami
>> neelkumar
```

- `which` - locates executable file

```bash
which java
>> /usr/bin/java
```

- `id` - tells info of users and groups

```bash
id
>> uid=501(neelkumar) gid=20(staff) groups=20(staff),12(everyone),61(localaccounts),79(_appserverusr),80(admin),81(_appserveradm),701(com.apple.sharepoint.group.1),33(_appstore),98(_lpadmin),100(_lpoperator),204(_developer),250(_analyticsusers),395(com.apple.access_ftp),398(com.apple.access_screensharing),399(com.apple.access_ssh),400(com.apple.access_remote_ae)
```

- `sudo` - “superuser do”

Allow regular (non root) user to run commands that needs system permission

Linux has two types of users:

1. **Normal user:** Limited permissions (can’t modify system files)
2. **Root (superuser):** Full control of the system

Without sudo, trying to restart os won’t work

```bash
shutdown -r now
>> shutdown: NOT super-user
```

With sudo, system will restart

```bash
sudo shutdown -r now
```

- `apt` - application package manager
1. making linux up to date

```bash
sudo apt-get update
```

1. Installing docker with apt

⇒ This won’t work as linux try to find [docker.io](http://docker.io) in our system which is not there

```bash
apt install docker.io
```

⇒ We use `apt-get` command

```bash
sudo apt-get install docker.io
```

1. finding where docker is installed

```bash
which docker
```

1. uninstalling docker

```bash
sudo apt remove docker
```

<hr>

### User and group management commands

- Adding user

```bash
sudo useradd -m krishna
```

Now go in home and check users,

```bash
cd ..
ls
>> krishna  neel-kumar
```

setting password for user krishna

```bash
sudo passwd krishna
```

switching to user krishna

```bash
su krishna
```

To go back / switch to primary user

```bash
exit
```

To remove krishna (user)

```bash
cd neel-kumar
sudo userdel krishna
su krishna
>> su: user krishna does not exist
```

- creating groups
1. first add few users then to check,

```bash
cat /etc/group
```

adding groups

```bash
sudo groupadd devops
sudo groupadd tester
```

adding users to groups

```bash
sudo gpasswd -a neel-kumar devops
sudo gpasswd -a ram devops
sudo gpasswd -a shyam devops
sudo gpasswd -a mohan devops
sudo gpasswd -a krishna tester
sudo gpasswd -a mohit tester
sudo gpasswd -a rohit tester
```

check them

```bash
cat /etc/group
>>devops:x:1007:neel-kumar,ram, shyam, mohan
>>devops:x:1008:krishna,mohit,rohit
```

adding multiple users

```bash
sudo gpasswd -M riya,siya devops
```

deleting group

```bash
sudo groupdel tester
```

<hr>

### File permission commands

Current folder structure:

```bash
/myFiles
			|_____ project.txt
			|_____ linuxDir
```

Run this: `ls -l`

```bash
ls -l

>> 
total 0
drwxr-xr-x  2 neelkumar  staff  64  7 Oct 13:51 linuxDir
-rw-r--r--  1 neelkumar  staff   0  7 Oct 13:49 project.txt
```

- `drwxr-xr-x` here `d` means directory
- then there is `rwxr-xr-x` - this is a triplet
    1. 1st triplet: `rwx` - this is USER
    2. 2nd triplet:  `r-x` - this is GROUP
    3. 3rd triplet:  `r-x` - this is OTHER

Here

`r` means read

`w` means write

`x` means execute

These are the permissions to user, group and other

`-`  means no none

| Value | read | write | execute |
| --- | --- | --- | --- |
| 0 | - | - | - |
| 1 | - | - | x |
| 2 | - | w | - |
| 3 | - | w | x |
| 4 | r | - | - |
| 5 | r | - | x |
| 6 | r | w | - |
| 7 | r | w | x |

so from this table we can say `rwx` has value `7`

so `rwxrwxr-x` means 775

we can use these vales to change file permissions

```bash
chmod 777 linuxDir
ls -l

>>
total 0
drwxrwxrwx  2 neelkumar  staff  64  7 Oct 13:51 linuxDir
-rw-r--r--  1 neelkumar  staff   0  7 Oct 13:49 project.txt
```

- umask

By default file system permission is known by `umask`

```bash
umask
>> 022
```

| umask | File perms | Dir perms | Description |
| --- | --- | --- | --- |
| 000 | 666 | 777 | Everyone full access |
| 022 | 644 | 755 | Standard default |
| 027 | 640 | 750 | Private group dirs |
| 077 | 600 | 700 | Very restricted (owner only) |
- changing ownership

```bash
ls -l

>>
total 4
drwxrwxr-x 2 neel-kumar neel-kumar 4096 Oct  7 15:27 linuxDir
-rw-rw-r-- 1 neel-kumar neel-kumar    0 Oct  7 15:27 project.txt
```

here both user and group is neel-kumar

to change user to krishna

```bash
sudo chown krishna project.txt
ls -l

>>
total 4
drwxrwxr-x 2 neel-kumar neel-kumar 4096 Oct  7 15:27 linuxDir
-rw-rw-r-- 1 krishna    neel-kumar    0 Oct  7 15:27 project.txt
```

- changing group

```bash
sudo chgrp devops linuxDir
ls -l

>>
total 4
drwxrwxr-x 2 neel-kumar devops     4096 Oct  7 15:27 linuxDir
-rw-rw-r-- 1 krishna    neel-kumar    0 Oct  7 15:27 project.txt
```

- To zip a folder `myFiles` that contains `linuxDir` and `project.txt`

```bash
zip -r myZip.zip myFiles
```

here `-r` means recursively - means to all zip all the files and folders inside `myFiles` otherwise an empty zip file will be created

- copying and unzipping zip file

```bash
mkdir zipFiles
cp myZip.zip zipFiles
cd zipFiles
unzip myzip.zip
```

- `tar` command

compressing folder

```bash
tar -cvzf folder1.tar.gz folder1

>>
a folder1
a folder1/file2.txt
a folder1/file3.txt
a folder1/file1.txt
```

⇒ a compressed `folder1.tar.gz` is created

copying it to other folder

```bash
cp folder1.tar.gz compressedFiles
```

now going in that folder and uncompressing file

```bash
cd compressedFiles
tar -xvzf folder1.tar.gz
```

Here in `xvzf` or `cvzf` means,

`x` - to extract

`c` - to compress

`v` - verbosely means to show details while executing

`z` - compress with gzip

`f` - file

<hr>

### File transfer commands

- sending file from local to remote server (linux vm box in our case)

Setups on mac terminal and linux vm box

1. On linux vm box

```bash
sudo apt update
sudo apt install openssh-server -y
```

1. check is ssh is active

```bash
sudo systemctl status ssh
```

1. If status: inactive

```bash
sudo systemctl start ssh
sudo systemctl enable ssh
```

1. in vm box go in network settings and choose network adapter `attached to: bridged adapter` and allow promiscuous mode
2. check ip address of linux

```bash
hostname -I
>>
192.168.29.165 2405:201:401e:c184:931d:743c:6cc4:f76c 2405:201:401e:c184:a00:27ff:fef4:b245 
```

here ip address : 192.168.29.165

This is used to connect with macos

1. check connection in mac terminal

```bash
ping 192.168.29.165
>>
64 bytes from 192.168.29.165: icmp_seq=1 ttl=64 time=6.288 ms
```

1. send file

```bash
scp ~/documents/workspace/linux/folder1/file1.txt neel-kumar@192.168.29.165:/home/neel-kumar/
```

- sending file from remote server to local
1. check mac ip address

```bash
ifconfig | grep "inet "
>>
inet 127.0.0.1 netmask 0xff000000
inet 192.168.29.77 netmask 0xffffff00 broadcast 192.168.29.255
```

here, ip address = `192.168.29.77`

1. send file

```bash
scp ~/test.txt neelkumar@192.168.29.77:/Users/neelkumar/Desktop/
```

<hr>

### Networking commands

- `ping` - send and receive data packets

```bash
ping google.com
>>
64 bytes from 142.251.223.206: icmp_seq=1 ttl=114 time=35.944 ms
```

tells that connection is working ✅

- `netstat`

```bash
sudo apt-get update
sudo apt install netstat
netstat
```

Tells about connections established

- `ifconfig` - shows all network interfaces

A network interface ****is basically a connection point between your computer and a network.

It’s what your system uses to send and receive data — like a “door” for network communication

```bash
ifconfig
```

- `traceroute`

tells you how your internet data travels ****from your computer to another place (like youtube.com).

It shows each stop (hop) ****your data makes along the way.

```bash
traceroute youtube.com
>>
 1  10.214.67.59 (10.214.67.59)  5.519 ms  7.374 ms  5.343 ms
 2  192.168.16.90 (192.168.16.90)  39.727 ms  206.339 ms  40.589 ms
 3  * * *
 4  * 192.168.19.5 (192.168.19.5)  66.762 ms  60.312 ms
 5  * * *
 6  * * *
 7  128.185.12.97 (128.185.12.97)  68.535 ms
    128.185.12.93 (128.185.12.93)  270.211 ms
    128.185.12.97 (128.185.12.97)  71.459 ms
 8  * * *
 9  * * *
10  172.253.67.98 (172.253.67.98)  107.597 ms
    142.251.52.206 (142.251.52.206)  49.577 ms
    172.253.73.194 (172.253.73.194)  28.085 ms
11  142.251.52.229 (142.251.52.229)  38.270 ms
    192.178.83.206 (192.178.83.206)  259.355 ms
    142.251.52.231 (142.251.52.231)  49.590 ms
12  del12s02-in-f14.1e100.net (142.250.194.46)  51.964 ms
    216.239.62.181 (216.239.62.181)  47.995 ms
    del12s02-in-f14.1e100.net (142.250.194.46)  60.987 ms
```

- `tracepath`

```bash
tracepath youtube.com
```

- `mtr` - my trace route

combo of `ping` + `traceroute`

It shows:

- The **path** your data takes (like traceroute)
- The **speed and packet loss** at each hop (like ping)
    
    👉 and it updates live — in real time!
    

```bash
mtr youtube.com
```

- `nslookup` - tells ip address of a website

```bash
nslookup youtube.com
>>
Server:		127.0.0.53
Address:	127.0.0.53#53

Non-authoritative answer:
Name:	youtube.com
Address: 172.217.26.110
Name:	youtube.com
Address: 2404:6800:4002:815::200e
```

- `telnet` - to connect to another server with specific port

```bash
telnet youtube.com 80
>>
Trying 2404:6800:4002:831::200e...
Connected to youtube.com
```

- `hostname` - tells name of host

```bash
hostname
```

- `ip address show` - shows our ip addresses of system

```bash
ip address show
```

- `iwconfig` - tells wireless connections

```bash
iwconfig
```

- `ss`  - socket statistics

similar to netstat

tell all connections and ports on system

```bash
ss
```

- `dig` - like nslookup

ask dns server for domain name and ip address

```bash
dig github.com
>>
; <<>> DiG 9.10.6 <<>> github.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 881
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 1280
;; QUESTION SECTION:
;github.com.			IN	A

;; ANSWER SECTION:
github.com.		54	IN	A	20.207.73.82

;; Query time: 77 msec
;; SERVER: 2401:4900:a06b:4897::4#53(2401:4900:a06b:4897::4)
;; WHEN: Wed Oct 08 06:20:38 IST 2025
;; MSG SIZE  rcvd: 55
```

- `whois` - tell who owns a domain name and other public info

```bash
whois github.com
```

- `arp` - address resolution protocol

tell which MAC address has which ip address on your network

```bash
arp -a
>>
? (10.214.67.26) at 8:0:27:f4:b2:45 on en0 ifscope [ethernet]
? (10.214.67.59) at 3e:44:7e:37:28:9 on en0 ifscope [ethernet]
```

gives ip:mac pair

- `ifplugstatus` - tells working status of interfaces

- Getting API data

```bash
curl -X GET https://jsonplaceholder.typicode.com/users
```

beautify the data you get from api

```bash
curl -X GET https://jsonplaceholder.typicode.com/users | jq
```

- downloading files from remote

```bash
mkdir myInstalls
cd myInstalls
wget https://file-examples.com/wp-content/storage/2017/10/file-sample_150kB.pdf
```

- `nmap` - network mapper

scans the network and tell

1. running services
2. connected devices
3. open ports

Check mac ip address to scan mac’s network

```bash
ipconfig ifaddr en0
>> 192.168.29.77
```

Now in linux vm box

```bash
nmap -v 192.168.29.77
>>
Starting Nmap 7.95 ( https://nmap.org ) at 2025-10-08 20:25 IST
Initiating Ping Scan at 20:25
Scanning 192.168.29.77 [2 ports]
Completed Ping Scan at 20:25, 0.01s elapsed (1 total hosts)
Initiating Parallel DNS resolution of 1 host. at 20:25
Completed Parallel DNS resolution of 1 host. at 20:25, 0.01s elapsed
Initiating Connect Scan at 20:25
Scanning 192.168.29.77 [1000 ports]
Discovered open port 5000/tcp on 192.168.29.77
RTTVAR has grown to over 2.3 seconds, decreasing to 2.0
RTTVAR has grown to over 2.3 seconds, decreasing to 2.0
Discovered open port 7000/tcp on 192.168.29.77
Increasing send delay for 192.168.29.77 from 0 to 5 due to max_successful_tryno increase to 4
Discovered open port 49152/tcp on 192.168.29.77
Completed Connect Scan at 20:25, 16.38s elapsed (1000 total ports)
Nmap scan report for 192.168.29.77
Host is up (0.0022s latency).
Not shown: 997 closed tcp ports (conn-refused)
PORT      STATE SERVICE
5000/tcp  open  upnp
7000/tcp  open  afs3-fileserver
49152/tcp open  unknown

Read data files from: /usr/bin/../share/nmap
Nmap done: 1 IP address (1 host up) scanned in 16.41 seconds
```

- `route` - gives the route or all gateways through which our request is sent to internet

```bash
route
```

<hr>

### Important linux commands

- `awk` - text processing tool

reads files line by line, split each line into fields and take actions on those fields

create file `app.log` and print its contents

```bash
awk '{print}' app.log
```

print 1st filed

```bash
awk '{print $1}' app.log
```

print 1st, 2nd field

```bash
awk '{print $1, $2}' app.log
```

3rd field mainly contains `info` or `debug`

we want only `debug`

```bash
awk '/DEBUG/ {print $1, $3, $4}' app.log
```

taking only filtered `DEBUG` content and copying to another log

```bash
awk '/DEBUG/ {print $1, $2, $3}' app.log > debug.log
```

finding total lines with debug

```bash
awk '/DEBUG/ {count++} END {print count}' app.log
>> 30
```

printing lines 2 to 5

```bash
awk 'NR >= 2 && NR <= 5 {print}' app.log
```

- awk is used when we have a formatted data

so we also use `sed`

printing only `INFO`

```bash
sed -n '/INFO/p' app.log
```

replacing DEBUG with DEBUGGING

```bash
sed 's/DEBUG/DEBUGGING/g' debug.log
```

to see at what lines INFO came

```bash
sed -n -e '/INFO/=' app.log
```

replace INFO to LOG in line 1 to 10

```bash
sed '1,10 s/INFO/LOG/g' app.log
```

now doing same but printing just line 1 to 10 and to quit at 11th line

```bash
sed '1,10 s/INFO/LOG/g; 1,0p; 11q' app.log
```

- `GREP`

global regular expression pattern

to match  any pattern in file

like finding all INFO

```bash
grep INFO app.log
```

to check both  small/upper case info

```bash
grep -i info app.log
```

to count info

```bash
grep -i -c info app.log
```
