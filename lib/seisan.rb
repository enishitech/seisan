require 'seisan/version'
require 'seisan/reporter'
require 'seisan/base_renderer'
require 'seisan/expense_renderer'
require 'gimlet'

module Seisan
  def self.default_dest_base_path
    File.join(Dir.pwd, 'output')
  end

  def self.report(config, dest_base_path=default_dest_base_path)
    src_path = File.join(Dir.pwd, 'data', config[:target])
    source = Gimlet::DataStore.new(src_path)

    requests = source.to_h.values
    exporter = Seisan::Reporter.new(requests, config)

    dest_path = File.join(dest_base_path, '%s.xlsx' % config[:target].gsub('/', '-'))
    exporter.export(dest_path)
  end
end
