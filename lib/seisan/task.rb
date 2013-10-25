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
        report(src_dir, dest_dir, config)
      end
      task :default => :seisan
    end

    private
    def report(src_dir, dest_dir, config)
      if config['target'].nil?
        Seisan.logger.error "You must specify the 'target'.\nExample:\n  % bundle exec rake target=2013/07"
        exit
      end

      Seisan.logger.info 'Processing %s ...' % config['target']
      requests = load_seisan_requests(src_dir, config['target'])
      Seisan.logger.info 'Loaded %d files' % requests.size
      report = Seisan::Report.new(requests, config)

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

    def convert_target_to_file_name(target)
      target.gsub('/', '-')
    end
  end
end

# Install tasks
Seisan::Task.install_tasks
