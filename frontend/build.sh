IMAGE='notes-frontend'

echo '>>> Building command...'
npm run build
echo '...Done'

echo '>>> Building image...'
docker image build --network host -t $IMAGE .
