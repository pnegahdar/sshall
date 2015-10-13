SSHAll 
=======

Install:

    go get github.com/pnegahdar/sshall/...
    go install github.com/pnegahdar/sshall

Usage:

    cat iplist.txt | sshall --concurrency 50 --cmd "ls -la"
     
     
 Recipes:
 
 Rotate Keys on AWS:
 
    aws ec2 describe-instances --query 'Reservations[].Instances[].[PrivateIpAddress]' --filters "Name=instance-state-name,Values=running" --output text \
    | sshall --concurrency 50 --try-user=ubuntu --try-user=ec2-user --cmd="echo $(cat ~/.ssh/key.pub) >> ~/.ssh/authorized_keys"        

