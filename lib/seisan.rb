require 'seisan/version'
require 'seisan/request'
require 'seisan/reporter'
require 'seisan/base_renderer'
require 'seisan/expense_renderer'

module Seisan
  def self.default_dest_base_path
    File.join(Dir.pwd, 'output')
  end

  def self.report(config, dest_base_path=default_dest_base_path)
    src_path = File.join(Dir.pwd, 'data', config[:target])
    request = Seisan::Request.new(src_path)
    request.export(config, dest_base_path)
  end
end
