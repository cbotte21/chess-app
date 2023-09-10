<#
	.SYNOPSIS Downloads source files of server(s), builds images and uploads to Packer.
#>

param (
	[string[]]$Servers = @([
		"chess-client-nextjs"
		"hive-go",
		"chess-go",
		"auth-go",
		"archive-go",
		"judicial-go"
	])
	[string]$InstallDir = "~/servers"
)

$Servers | For-EachObject {
	mkdir ~/servers
	git clone $_ $InstallDir
	# Build go projects
	# Publish to PACKERHUB
}
