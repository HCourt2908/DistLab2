# Distributed Lab 2: RPC in Go

## Using the lab sheet

There are two ways to use the lab sheet, you can either:

- [create a new repo from this template](https://github.com/UoB-CSA/distributed-lab-2/generate) - **this is the recommended way**
- download a [zip file](https://github.com/UoB-CSA/distributed-lab-2/archive/master.zip)

## Ratings

Each question is rated to help you balance your work:

- :red_circle::white_circle::white_circle::white_circle::white_circle: - Easy, strictly necessary.
- :red_circle::red_circle::white_circle::white_circle::white_circle: - Medium, still necessary.
- :red_circle::red_circle::red_circle::white_circle::white_circle: - Hard, necessary if you're aiming for higher marks.
- :red_circle::red_circle::red_circle::red_circle::white_circle: - Hard, useful for coursework extensions.
- :red_circle::red_circle::red_circle::red_circle::red_circle: - Hard, beyond what you need for any part of the coursework.

## 1: Deploying the Chat System :red_circle::white_circle::white_circle::white_circle::white_circle:

Following from the end of last week's lab, you should now have a distributed
system -- a chat client and server. Hopefully you've already tested these out 
locally, but now you need to put them to the test by running them in a properly
distributed fashion.

For this, we will make use of AWS. Follow the guides for setting
up instances, accessing them, and opening required ports.

In particular, you should:

+ create two t2.micro instances, ([guide](https://github.com/UoB-CSA/setup-guides/blob/master/aws/create-instance.md))
+ make sure Go is installed on your instances ([guide](https://github.com/UoB-CSA/setup-guides/blob/master/go-install/aws.md#setup-instructions-for-aws-ubuntu-2004-lts-focal))
+ load your client/server code onto the instances (e.g., via [git](https://github.com/UoB-CSA/setup-guides/blob/master/aws/access-instance.md#coping-files))
+ open ports to allow communication between instances ([guide](https://github.com/UoB-CSA/setup-guides/blob/master/aws/ports.md))
+ start your server running on one instance ([ssh guide](https://github.com/UoB-CSA/setup-guides/blob/master/aws/access-instance.md))
+ connect to your server with your client on another instance
+ connect to your server from your local machine
+ communicate between clients via your (genuinely) distributed system!

**Make sure you stop your instances when you have finished!**

### Frequently Asked Questions

- **My code worked locally but it doesn't on AWS:** The main reason for this happening is incorrect IP addresses. To connect to the server from another AWS instance, you will need to use the private IP address. To connect to the server from a non-AWS computer (i.e. your personal computer) you need to use the public IP address. Think about why this may be the case, how do private and public IP addresses work?

- **I use `go run client.go` and get `panic: runtime error: invalid memory address or nil pointer dereference`:** You are not passing in the IP address of the server to the client. Run it like this `go run client.go -ip ip:port` replacing `ip` and `port` with the appropriate values.

- **What is git? How do I copy files with it?** [Git](https://cs-uob.github.io/COMS10012/exercises/part1/git/index.html) is a version control system, it allows you to keep track of your changes. When we suggest loading files to AWS, we are suggesting that you should use a git repository and pushing your changes to the remote repository, then cloning/pulling these changes to your AWS instance. GitHub, GitLab and many others are services that allow for this process. Still stuck? You can try using [SCP](https://linuxize.com/post/how-to-use-scp-command-to-securely-transfer-files/) - we do not recommend this though. You should be familiar with git at this point.
  
- Alternatively, **for copying files** consider using scp. The command to copy `my_stuff.zip` into the home directory on your EC2 instance will be something like `scp -i ~/.ssh/keypair.pem my_stuff.zip ec2-user@54.210.24.40:~`. Replace the IP address with your IP and `my_stuff.zip` with the file you want to copy. This method is a frustrating as you will need to zip and unzip stuff all the time.


- Another option is to use a GUI based FTP program. (WinSCP)[https://winscp.net/eng/index.php] is a great option for Windows users. I cannot recommend FileZilla in good faith FileZilla as **some vendors bundle it with Malware!!** However, I assume if you install with apt from the Ubuntu repos it will be ok (have not tested this!)


## 2: Using RPC - Secret Strings :red_circle::white_circle::white_circle::white_circle::white_circle:

Follow the video from this week to create a simple RPC system that allows
clients to call the functions of a server via a defined interface. As with last
week, this is best attempted in stages:

+ Stage 1: Write server code to enable access to the "secret" string
manipulation function. Test it by writing a client that sends a string to be
reversed.
+ Stage 2: Enable Premium Tier service by implementing the `FastReverse`
function in the server.
+ Stage 3: Update your client to read words from the `wordlist` file and reverse them all.
+ (Optional) run multiple server instances, and speed up the work by
having your client split the load between servers.

## 3: 99 Bottles of Beer :red_circle::red_circle::red_circle::white_circle::white_circle:

So far we've been focusing on client-server systems. However, there are times
when we don't really want a distinction in roles between components -- when we
might want them to act as peers, all running the same code.

For this task, you are going to solve the ["99 Bottles of
Beer"](https://en.wikipedia.org/wiki/99_Bottles_of_Beer) problem in a
distributed fashion. Rather than a single process singing the lyrics all by itself,
you're going to run at least three instances of your code, on different
machines, and have these buddies share the work of singing the song
verse-by-verse, in order. 

Here's how it should go:

+ **Buddy 1**: "99 bottles of beer on the wall, 99 bottles of beer. Take one down, pass it around..."
+ **Buddy 2**: "98 bottles of beer on the wall, 98 bottles of beer. Take one down, pass it around..."
+ **Buddy 3**: "97 bottles of beer on the wall, 97 bottles of beer. Take one down, pass it around..."
+ **Buddy 1**: "96 bottles of beer on the wall, 96 bottles of beer. Take one down, pass it around..."

All the way down until there are no more bottles of beer on the wall, which the
final singer should note appropriately with your preferred ending to the song.

Use 3 for testing, but your solution should be able to handle any number of
buddies singing along. There should be no difference in the code, just different
flags on the commandline.

### Hints: 

1. Each process needs to accept an `ip:port` string for the 'next' buddy who will
follow on from them in the song. You'll have to configure them in a loop. 
2. You don't want clients to try connect to each other straight away, or you
won't have time to set the final process running so that the first can connect.
3. When you set up the processes, you'll also need some way to indicate which of them 
should start the song (I suggest allowing any `n` bottles of beer, for testing purposes). 
Only the last process you set up should need to be told the `n` to count down from.
4. You may need to look at `client.Go` rather than `client.Call`.

## Don't forget to stop any AWS instances!
