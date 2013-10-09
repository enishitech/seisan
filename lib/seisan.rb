require 'seisan/version'
require 'seisan/report'
require 'seisan/base_renderer'
require 'seisan/header_renderer'
require 'seisan/expense_renderer'
require 'logger'

module Seisan
  def self.logger
    @@logger ||= initialize_default_logger
  end

  def self.logger=(logger)
    @@logger = logger
  end

  def self.initialize_default_logger
    logger = Logger.new(STDOUT)
    logger.formatter = proc{|severity, datetime, progname, message|
      "#{message}\n"
    }
    logger
  end
end
