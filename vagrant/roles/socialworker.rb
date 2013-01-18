name "socialworker"
description "The  role for socialWorker servers"

run_list ["recipe[nodejs]","recipe[golang]"]

default_attributes({ 
                     "kd_deploy" => {
                                "revision_tag" => "HEAD",
                                "release_action" => :deploy,
                                "deploy_dir" => '/opt/koding',
                     },
                     "launch" => {
                                "config" => "autoscale",
                                "programs" => ["socialWorker"]
                     }
})
