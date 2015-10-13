SSHAll 
=======

Install:

    go install github.com/pnegahdar/sshall/...

Usage:

    cat iplist.txt | sshall --concurrency 50 --cmd "ls -la"
     
     
 Recipes:
 
 Add key to AWS boxes:
 
    aws ec2 describe-instances --query 'Reservations[].Instances[].[PrivateIpAddress]' --output text | grep -v None | sshall --concurrency 50 --cmd "ls"        
