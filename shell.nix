{ pkgs }:
pkgs.mkShell {
	name = "auto-brightness-devshell";

	buildInputs = [
		pkgs.go
	];
}
