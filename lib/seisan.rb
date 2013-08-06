require 'seisan/version'
require 'seisan/request'
require 'seisan/reporter'
require 'seisan/base_renderer'
require 'seisan/expense_renderer'

module Seisan
  def self.src_path(year, month)
    File.join(Dir.pwd, 'data', '%04d' % year, '%02d' % month)
  end

  def self.default_dest_base_path
    File.join(Dir.pwd, 'output')
  end

  def self.report(config, dest_base_path=default_dest_base_path)
    now = Time.now
    month = (config[:month] || now.month).to_i
    year = (config[:year] || now.year).to_i
    unless (1..12).include?(month)
        raise ArgumentError, "illigal month specified"
    end

    request = Seisan::Request.new(src_path(year, month))
    request.export(config, dest_base_path)
  end
end
