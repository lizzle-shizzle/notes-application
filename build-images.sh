echo '========= Building Images'
cd backend && ./build.sh && cd ../
cd frontend && ./build.sh && cd ../
echo '========= Build Complete.'

echo '========= Updating Host Names'
echo "127.0.0.1	www.notes-application.co.za" >> /etc/hosts
