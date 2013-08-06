require 'axlsx'
require 'fileutils'
require 'seisan/expense_renderer'

module Seisan
  class Reporter
    RENDERES = [Seisan::ExpenseRenderer]

    def initialize(request, config)
      @request = request
      @config = config
    end

    def export(dest_path)
      prepare_sheet

      render_global_header

      section_renderers.each do |renderer|
        row
        renderer.render
      end

      write_to_file(dest_path)
    end

    private
    def section_renderers
      @section_renderers ||= RENDERES.map {|r| r.new(@request, &self.method('row')) }
    end

    def render_global_header
      row ["#{@config[:organization_name]} 精算シート %04d年%02d月" % [@config[:year], @config[:month]]]
      row ['作成時刻', Time.now.strftime('%Y-%m-%d %X')]
    end

    def row(columns=[])
      @sheet.add_row columns, :style => @font
    end

    def prepare_sheet
      @package = Axlsx::Package.new
      @workbook = @package.workbook
      @package.use_shared_strings = true
      @font = @workbook.styles.add_style :font_name => 'ＭＳ Ｐゴシック'
      @sheet = @workbook.add_worksheet(:name => '精算シート')
    end

    def write_to_file(dest_path)
      FileUtils.mkdir_p(File.dirname(dest_path))
      @package.serialize(dest_path)
      puts 'Wrote to %s' % dest_path
    end
  end
end
