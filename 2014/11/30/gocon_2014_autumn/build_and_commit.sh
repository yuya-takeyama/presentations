orig_path=`pwd`
sh build.sh
git checkout gh-pages
mv slide.html index.html
git add index.html
git commit
git checkout master
cd $orig_path
