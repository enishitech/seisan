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
        config = user_config.merge('target' => ENV['target'])
        report(src_dir, dest_dir, config, $stdout)
      end
      task :default => :seisan
    end

    private
    def report(src_dir, dest_dir, config, output)
      requests = load_seisan_requests(src_dir, config['target'])
      display_load_status(requests, output)
      report = Seisan::Report.new(requests, config, output)

      dest_path = File.join(dest_dir, '%s.xlsx' % convert_target_to_file_name(config['target']))
      report.export(dest_path)
    end

    def user_config
      Gimlet::DataStore.new('config.yaml').to_h
    rescue Gimlet::DataStore::SourceNotFound
      {}
    end

    def load_seisan_requests(src_dir, target)
      source = Gimlet::DataStore.new(File.join(src_dir, target))
      source.to_h.values
    rescue Gimlet::DataStore::SourceNotFound
      []
    end

    def display_load_status(requests, output)
      entries = requests.empty? ? [] : requests.map {|e| e[:expense] }.flatten
      output.puts 'Loaded %d files, %d expense entries.' % [requests.size, entries.size]
    end

    def convert_target_to_file_name(target)
      target.gsub('/', '-')
    end
  end
end

# Install tasks
Seisan::Task.install_tasks
