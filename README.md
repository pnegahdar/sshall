SSHAll 
=======

Install:

    go install github.com/pnegahdar/sshall/...

Usage:

    cat iplist.txt | sshall "ls -la"
     
     
 Recipes:
 
 Add key to AWS boxes:
 
    aws ec2 describe-instances --query 'Reservations[].Instances[].[PrivateIpAddress]' --output text | sshall -i ~/.ssh/id_old "echo $(~/.ssh/id_rsa.pub)" > ~/.ssh/authorized_keys"
        
