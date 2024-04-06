{
	description = "Automatically adjust brightness over time";

	inputs = {
		nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
		flake-utils.url = "github:numtide/flake-utils";
	};

	outputs = { self, nixpkgs, flake-utils }:
		flake-utils.lib.eachDefaultSystem
			(system:
				let pkgs = import nixpkgs {
					system = system;
				}; in
				{
					defaultPackage = pkgs.buildGoModule {
						pname = "auto_brightness";
						version = "0.1.0";
						src = ./.;
						#vendorHash = pkgs.lib.fakeHash;
						vendorHash = null;
					};
				}
			);
}
