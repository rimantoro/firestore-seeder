#!/bin/bash

collections=( "distributor_permissions" "distributor_roles" "distributor_users" "distributors" "distributor_vaccines" "medfac_permissions" "medfac_roles" "medfac_users" "medfacs" "ops_permissions" "ops_roles" "ops_users" )

for coll in "${collections[@]}"
do
   firebase firestore:delete --non-interactive -ry --project spartech-282005 ${coll}
  # echo ${coll}
done
