BUCKET="$1"
exists=$(aws s3 ls s3://$BUCKET/terraform.tfstate)
if [ -z "$exists" ]; then
  terraform init
  terraform apply -auto-approve
  aws s3 cp terraform.tfstate s3://$BUCKET/
else
  echo "terraform.tfstate is allready exists"
fi


function reapply {
    exists=$(aws s3 ls s3://$BUCKET/terraform.tfstate)
    if [ -z "$exists" ]; then
      terraform init
      terraform apply -auto-approve
      aws s3 cp terraform.tfstate s3://$BUCKET/
    else
      aws s3 cp s3://$BUCKET/terraform.tfstate .
      terraform init
      terraform apply -auto-approve
      aws s3 cp terraform.tfstate s3://$BUCKET/
    fi
}
function apply {
	  terraform init
      terraform apply -auto-approve
      aws s3 cp terraform.tfstate s3://$BUCKET/
  }


function help {
    echo "The script is designed to automate file storage and use of terraform"
    echo "it checks that the file is in your storage and writes/overwrites" 
    echo "it if necessary"
    echo
    echo "Directory structure for script work:"
    echo '/$BUCKET                      (name of your storage s3 bucket)'
    echo "|-tftry.sh                    (this file)"
    echo "|-main.tf                     (terraform init )"
    echo "|-/terraform-tokend-vanilla   (folder you initscript)" 
    echo
    echo "Syntax: scriptTemplate [-apply|reapply]"
    echo "options:"
    echo 
    echo "reapply     updates the terraform.tfstate file in the storage"
    echo "apply       makes new terraform.tfstate file in the storage"
}

echo
while [ -n "$1" ]
do
case "$1" in
-h)
help;;

-apply)
apply $1;;

-reapply) 
reapply $1;;
esac
shift
done 
