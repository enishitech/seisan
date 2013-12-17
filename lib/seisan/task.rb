require 'rake'
require 'seisan'
require 'gimlet'

module Seisan
  class Task
    include Rake::DSL if defined? Rake::DSL
    def self.install_tasks
      new.install
    end

    def install
      desc "Generate seisan report"
        config = {
          organization_name: '株式会社えにしテック',
          src_base_path: default_src_base_path,
          dest_base_path: default_dest_base_path,
          target: ENV['target'],
        }
        report(config)
      task :seisan do
      end
      task :default => :seisan
    end

    def report(config)
      source = Gimlet::DataStore.new(File.join(config[:src_base_path], config[:target]))

      requests = source.to_h.values
      report = Seisan::Report.new(requests, config)

      dest_path = File.join(config[:dest_base_path], '%s.xlsx' % config[:target].gsub('/', '-'))
      report.export(dest_path)
    end

    def default_src_base_path
      File.join(Dir.pwd, 'data')
    end

    def default_dest_base_path
      File.join(Dir.pwd, 'output')
    end
  end
end

# Install tasks
Seisan::Task.install_tasks
