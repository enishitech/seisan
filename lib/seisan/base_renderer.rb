module Seisan
  class BaseRenderer
    def initialize(requests, sheet, font, config)
      @requests = requests
      @sheet = sheet
      @font = font
      @config = config
    end

    def requests
      @requests
    end

    def row(columns=[])
      @sheet.add_row columns, :style => @font
    end
  end
end
