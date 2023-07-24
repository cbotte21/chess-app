param (
	[Parameter(mandatory=$true)]
	[ValidateSet("docker-compose","kubernetes")]
	[string]$Deployment = "docker-compose"
)

Set-Location "$PSScriptRoot/../infrastructure/docker-compose"

$Flags = $null
switch ($Deployment) {
	"docker-compose" {
		$Flags = @{
			FilePath = "docker"
			ArgumentList = "compose build"
		}
	}
	default {
		Write-Error "Not yet implemented."
	}
}

Start-Process @Flags
