module Seisan
  class BaseRenderer
    def initialize(request, &render_row_method)
      @request = request
      @render_row_method = render_row_method
    end

    def request
      @request
    end

    def row
      @render_row_method
    end
  end
end
