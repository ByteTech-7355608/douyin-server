#! bin/bash
if [ $2=="start" ];then
echo "begin~"
sh ./run_api.sh >api.log  &
echo "$!" > pid
sh ./run_base.sh >base.log  & 
echo "$!" > pid
sh ./run_interaction.sh >interaction.log  & 
echo "$!" > pid
sh ./run_social.sh >social.log  &
echo "$!" > pid
sh ./run_cron.sh >cron.log  &
echo "$!" > pid
elif [ $2=="stop" ];then
    kill `cat pid`
    echo "finished~"
else
echo "input correct condition"
fi
