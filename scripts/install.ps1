param (
	[parameter(Mandatory)]
	[string]$ConfigFile,
	[string[]]$TargetServers = @()
)

function Get-Addr {
	param(
		[parameter(Mandatory)]
		[string]$Name
	)
	"$($Name.Split('-')[$Name.Split('-').Length - 2]).default.svc.cluster.local"
}

# Load and parse server config file

$Config = Get-Content -Raw -Path $ConfigFile | ConvertFrom-Json
$Servers = $Config.servers

# Actual script

Push-Location "$PSScriptRoot/../infrastructure/packer"

$Jobs = @()

$Servers | ForEach-Object {
	# If servers are targeted, skip untargeted
	if ($TargetServers.Count -gt 0 -and -not ($TargetServers -contains $_.name)) {
		return
	}

	[string]$commitSha = (git ls-remote "https://github.com/cbotte21/$($_.name)").Replace("`t", " ").Split(' ')[0]
	[string[]]$tmpNameSplit = $_.name.Split('-')
	[string]$serverType = $tmpNameSplit[$tmpNameSplit.Length-1] # Taken by extension ex) server-go -> go

	# Parse environment variables
	[string]$setEnvVars = ""
	$_.environment_variables | ForEach-Object {
		$setEnvVars += "'$_' "
	}
	$setEnvVars += "'port=$($_.port)'"

	# Parse server links
	$_.server_links | ForEach-Object {
		$s = $_.Split('-')
		$name = $_.Substring(0, $_.Length-$s[$s.length-1].Length-1)
		$setEnvVars += "'$($name)_addr=$(Get-Addr $_)' "
	}

	# Set environment variables for packer
	$env:PKR_VAR_name = $_.name
	$env:PKR_VAR_version = $commitSha
	$env:PKR_VAR_port = $_.port
	$env:PKR_VAR_set_environment = $setEnvVars

	# Start job to build image
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
