class Microservice {
	[string]$Name
	[string]$Port
	[String[]]$EnvironmentVariables

	Microservice([string]$Name, [string]$Port, [string[]]$EnvironmentVariables) {
		$this.Name = $Name
		$this.Port = $Port
		$this.EnvironmentVariables = $EnvironmentVariables
	}
}

function Get-Addr {
	param(
		[parameter(Mandatory)]
		[string]$Name
	)
	"$($Name.Split('-')[$Name.Split('-').Length - 2]).default.svc.cluster.local"
}

$JWT_SECRET = "MYSUPERSECRETPASSCODE"
$MONGOURI = "mongodb://root:example@mongo.default.svc.cluster.local:27017/admin"

[Microservice[]]$Servers = @(
	[Microservice]::new(
			"auth-go",
			"5000",
			@(
				"jwt_secret=$JWT_SECRET",
				"mongo_uri=$MONGOURI"
				"mongo_db=auth"
			)
	),
	[Microservice]::new(
			"archive-go",
			"5001",
			@(
				"mongo_uri=$MONGOURI",
				"mongo_db=archive"
			)
	),
	[Microservice]::new(
			"chess-go",
			"5002",
			@(
				"queue_addr=$(Get-Addr "queue-go")",
				"jwt_secret=$JWT_SECRET",
				"redis_addr=$(Get-Addr "redis-server")"
			)
	)
	[Microservice]::new(
			"queue-go",
			"5003",
			@(
				"chess_addr=$(Get-Addr "chess-go")",
				"redis_addr=$(Get-Addr "redis-server")"
			)
	),
	[Microservice]::new(
			"hive-go",
			"5004",
			@(
				"jwt_secret=$JWT_SECRET",
				"judicial_addr=$(Get-Addr "judicial-go")",
				"redis_addr=$(Get-Addr "redis-server")"
			)
	),
	[Microservice]::new(
			"username-go",
			"5005",
			@(
				"mongo_uri=$MONGOURI",
				"jwt_secret=$JWT_SECRET"
				"mongo_db=username"
			)
	),
	[Microservice]::new(
			"judicial-go",
			"5006",
			@(
				"mongo_uri=$MONGOURI",
				"hive_addr=$(Get-Addr "hive-go")"
			)
	),
	[Microservice]::new(
			"chess-client-nextjs",
			"3000",
			@(
				"hive_addr=$(Get-Addr "hive-go")",
				"queue_addr=$(Get-Addr "queue-go")",
				"proto_dir=chess-app/proto"
			)
	)
)

Push-Location "$PSScriptRoot/../infrastructure/packer"

$Jobs = @()
for ($index = 0; $index -lt $Servers.length; $index++) {
	[string]$serverName = $Servers[$index].Name
	[string]$version = (git ls-remote "https://github.com/cbotte21/${serverName}").Replace("`t", " ").Split(' ')[0]
	[string]$serverType = $serverName.Split('-')[$serverName.Split('-').Length-1]
	[string]$setEnvVars = ""

	$Servers[$index].EnvironmentVariables | ForEach-Object {
		$setEnvVars += "'$_' "
    }
	$setEnvVars += "'port=$($Servers[$index].Port)'"

	$env:PKR_VAR_name = $serverName
	$env:PKR_VAR_version = $version
	$env:PKR_VAR_port = $Servers[$index].Port
	$env:PKR_VAR_set_environment = $setEnvVars

	$Jobs += Start-ThreadJob -ScriptBlock {
		$processOptions = @{
			FilePath = "packer"
			ArgumentList = @(
			 	"build",
				"$using:serverType.pkr.hcl"
			)
		}
		Start-Process @processOptions
	}
}

Pop-Location
$Jobs | Receive-Job -Wait -AutoRemoveJob
