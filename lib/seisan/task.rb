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
      task :seisan do
        src_dir, dest_dir = 'data', 'output'
        report(src_dir, dest_dir, ENV['target'], $stdout)
      end
      task :default => :seisan
    end

    private
    def report(src_dir, dest_dir, target, output)
      requests = load_seisan_requests(src_dir, target)
      report = Seisan::Report.new(requests, target, user_config, output)

      dest_path = File.join(dest_dir, '%s.xlsx' % convert_target_to_file_name(target))
      report.export(dest_path)
    end

    def user_config
      Gimlet::DataStore.new('config.yaml')
    rescue Gimlet::DataStore::SourceNotFound
      {}
    end

    def load_seisan_requests(src_dir, target)
      source = Gimlet::DataStore.new(File.join(src_dir, target))
      source.to_h.values
    rescue Gimlet::DataStore::SourceNotFound
      []
    end

    def convert_target_to_file_name(target)
      target.gsub('/', '-')
    end
  end
end

# Install tasks
Seisan::Task.install_tasks
