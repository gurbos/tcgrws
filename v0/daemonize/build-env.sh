#! /bin/bash

# build-env.sh populates the environment.conf file with data
# needed to configure the web service.


CONFIG_FILE=./config-data.conf

NL=$'\n' # Newline character

source ./utils.sh

# Public host DNS name
PUBLIC_HOSTNAME=`/opt/aws/bin/ec2-metadata -p | cut -f2 -d:`
echo "PUBLIC_HOSTNAME=$PUBLIC_HOSTNAME" > $CONFIG_FILE

# Public RDS host name
DB_HOST=`aws rds describe-db-instances --query "DBInstances[?DBInstanceIdentifier=='webservicedb'].Endpoint.Address" --output text`
echo "DB_HOST=$DB_HOST" >> $CONFIG_FILE

DB_PORT=`aws rds describe-db-instances --query "DBInstances[?DBInstanceIdentifier=='webservicedb'].Endpoint.Port" --output text`
echo "DB_PORT=$DB_PORT" >> $CONFIG_FILE

# STATIC_DATA_PATH=$(staticDataDir)

# if [ "$STATIC_DATA_PATH" == "" ]; then
#     echo "ERROR: AWS EFS filesystem not found"
#     exit 1
# fi
