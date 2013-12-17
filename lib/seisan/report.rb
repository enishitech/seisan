require 'axlsx'
require 'fileutils'
require 'seisan/expense_renderer'

module Seisan
  class Report
    RENDERES = [Seisan::ExpenseRenderer]

    def initialize(requests, target, config, output)
      @requests = requests
      @target = target
      @config = config
      @output = output
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
      @section_renderers ||= RENDERES.map {|r| r.new(@requests, &self.method('row')) }
    end

    def render_global_header
      row ["#{organization_name} 精算シート #{target_name}"]
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
      @output.puts 'Wrote to %s' % dest_path
    end

    def target_name
      @target
    end

    def organization_name
      @config[:organization] ? @config[:organization][:name] : ''
    end
  end
end
