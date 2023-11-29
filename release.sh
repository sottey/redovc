pushd .
cd ~/go/src/projects/redo.vc
go build -o redovc
rake build
cd redovc_docs
mkdocs build
cd site
zip -r redovc_docs.zip * 
mv ./redovc_docs.zip ~/Downloads
popd
