require 'gimlet'

module Seisan
  class Request
    def initialize(src_path)
      @source = Gimlet::DataStore.new(src_path)
    end

    def export(config, dest_base_path)
      exporter = Seisan::Reporter.new(self, config)
      dest_path = File.join(dest_base_path, '%04d-%02d.xlsx' % [config[:year], config[:month]])
      exporter.export(dest_path)
    end

    def entries
      @source.to_h.values
    end
  end
end
