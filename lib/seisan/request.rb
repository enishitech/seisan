require 'gimlet'

module Seisan
  class Request
    def initialize(src_path)
      @source = Gimlet::DataStore.new(src_path)
    end

    def export(config, dest_base_path)
      exporter = Seisan::Reporter.new(entries, config)
      dest_path = File.join(dest_base_path, '%s.xlsx' % config[:target].gsub('/', '-'))
      exporter.export(dest_path)
    end

    def entries
      @source.to_h.values
    end
  end
end
