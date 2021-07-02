#! /bin/bash

# build-env.sh populates the environment.conf file with data
# needed to configure the web service.


CONFIG_FILE=./config-data.conf

ENV_VARS=""

NL=$'\n' # Newline character

source ./utils.sh

# Public host DNS name
# PUBLIC_HOSTNAME=`/opt/aws/bin/ec2-metadata -p | cut -f2 -d:`
# ENV_VARS+="PUBLIC_HOSTNAME=$PUBLIC_HOSTNAME"

# Database connection data (package tcgrws/dbio configuration data)
# DB_HOST=`aws rds describe-db-instances --query "DBInstances[?DBInstanceIdentifier=='webservicedb'].Endpoint.Address" --output text`
# DB_PORT=`aws rds describe-db-instances --query "DBInstances[?DBInstanceIdentifier=='webservicedb'].Endpoint.Port" --output text`
ENV_VARS="$ENV_VARS DB_HOST=$DB_HOST"
ENV_VARS="\n$ENV_VARS DB_PORT=$DB_PORT"
ENV_VARS="\n$ENV_VARS DB_USER=$DB_USER"
ENV_VARS="\n$ENV_VARS DB_PASSWD=$nDB_PASSWD"
ENV_VARS="\n$ENV_VARS DB_NAME=$DB_NAME"


# Static content location (package tcgrws/handlers configuration data)
ENV_VARS="\n$ENV_VARS HOST=$HOST"
ENV_VARS="\n$ENV_VARS STATIC_DATA_DIR=$STATIC_DATA_DIR"
# STATIC_DATA_DIR=$(FindStaticDataDir)
# if [ "$STATIC_DATA_PATH" == "" ]; then
#     echo "ERROR: AWS EFS filesystem not found"
#     exit 1
# fi
