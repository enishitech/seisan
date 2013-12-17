module Seisan
  class BaseRenderer
    def initialize(requests, &render_row_method)
      @requests = requests
      @render_row_method = render_row_method
    end

    def requests
      @requests
    end

    def row
      @render_row_method
    end
  end
end
