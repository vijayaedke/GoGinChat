#!/bin/sh

DBNAME = $1
DBUSER = $2
DBPWD  = $3
DBURL  = $4
COLLECTION_NAME = "users"

echo "Add Default admin user to the system"

echo "=============================================================================="
echo "====================Checking Mongo Configurations============================="
echo "=============================================================================="

echo `mongo --version`
echo `mongo --eval "printjson(db.serverStatus()) --quiet" `

createDBUser = `mongo --eval "printjson(db.createUser({user: $DBUSER, pwd: $DBPWD, roles: []}))"`
if [ $createDBUser -ne 0 ]; then
    echo "$DBUSER already exists"
    exit(0)
else
    echo "$DBUSER added"
    echo "=============================================================================="
    echo "=============================Adding Mongo User================================"
    echo "=============================================================================="

    echo `mongo --eval $DBURL/$DBNAME 'var document = { username : "admin", role : ["system_admin"] }; db.${COLLECTION_NAME}.insert(document);'`
fi