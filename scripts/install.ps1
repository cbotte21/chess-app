class Microservice {
	[string]$Name
	[string]$Port
	[String]$EnvironmentVariables

	Microservice([string]$Name, [string]$Port, [string]$EnvironmentVariables) {
		$this.Name = $Name
		$this.Port = $Port
		$this.EnvironmentVariables = $EnvironmentVariables
	}
}

$JWT_SECRET = "MYSUPERSECRETPASSCODE"
$MONGOURI = ""

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
			"archive-go", # Not compiling
			"5001",
			@(
				"mongo_uri=$MONGOURI",
				"mongo_db=archive"
			)
	),
	[Microservice]::new(
			"chess-go", # Not compiling
			"5002",
			@(
				"queue_addr=",
				"jwt_secret=$JWT_SECRET",
				"redis_addr="
			)
	),
	[Microservice]::new(
			"queue-go",
			"5003",
			@(
				"chess_addr=",
				"redis_addr="
			)
	),
	[Microservice]::new(
			"hive-go",
			"5004",
			@(
				"jwt_secret=$JWT_SECRET",
				"judicial_addr=",
				"redis_addr="
			)
	),
	[Microservice]::new(
			"username-go",
			"5005",
			@(
				"mongo_uri=$MONGOURI",
				"monogo_db=username"
			)
	),
	[Microservice]::new(
			"chess-client-nextjs",
			"3000",
			@()
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
		$setEnvVars += "echo '$_' >> .env; "
    }
	$setEnvVars += "echo 'port:$($Servers[$index].Port)' >> .env; "

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