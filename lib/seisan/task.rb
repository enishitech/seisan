require 'rake'
require 'seisan'

module Seisan
  class Task
    include Rake::DSL if defined? Rake::DSL
    def self.install_tasks
      new.install
    end

    def install
      desc "Generate seisan report"
      task :report do
        config = {
          organization_name: '株式会社えにしテック',
          target: ENV['target'],
        }
        Seisan.report(config)
      end
      task :default => :report
    end
  end
end

# Install tasks
Seisan::Task.install_tasks
