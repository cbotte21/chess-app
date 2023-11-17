param (
	[Parameter(mandatory=$true)]
	[ValidateSet("docker-compose","kubernetes")]
	[string]$Deployment = "docker-compose"
)

Push-Location "$PSScriptRoot/../infrastructure/docker-compose"

$Flags = $null
switch ($Deployment) {
	"docker-compose" {
		$Flags = @{
			FilePath = "docker"
			ArgumentList = "compose up"
		}
	}
	default {
		Write-Error "Not yet implemented."
	}
}

Pop-Location

Start-Process @Flags
