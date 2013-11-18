# coding: utf-8
lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'seisan/version'

Gem::Specification.new do |spec|
  spec.name          = "seisan"
  spec.version       = Seisan::VERSION
  spec.authors       = ["SHIMADA Koji"]
  spec.email         = ["koji.shimada@enishi-tech.com"]
  spec.description   = %q{seisan solution for small team}
  spec.summary       = %q{seisan solution for small team}
  spec.homepage      = ""
  spec.license       = "MIT"

  spec.files         = `git ls-files`.split($/)
  spec.executables   = spec.files.grep(%r{^bin/}) { |f| File.basename(f) }
  spec.test_files    = spec.files.grep(%r{^(test|spec|features)/})
  spec.require_paths = ["lib"]

  spec.add_runtime_dependency "axlsx", "~> 2.0.1"
  spec.add_runtime_dependency "gimlet", "~> 0.0.3"

  spec.add_development_dependency "bundler", "~> 1.3"
  spec.add_development_dependency "rake"
end
