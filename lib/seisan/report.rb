require 'axlsx'
require 'fileutils'
require 'seisan/header_renderer'
require 'seisan/expense_renderer'

module Seisan
  class Report
    DEFAULT_RENDERERS = [
      Seisan::HeaderRenderer,
      Seisan::ExpenseRenderer
    ]
    @@renderers = DEFAULT_RENDERERS

    class << self
      def renderer_chain(&block)
        @@renderers = []
        block.call(self) if block
      end

      def add(renderer)
        @@renderers << renderer
      end
    end

    def initialize(requests, config, output)
      @requests = requests
      @config = config
      @output = output
    end

    def export(dest_path)
      prepare_sheet

      renderers.each do |renderer|
        renderer.render
      end

      write_to_file(dest_path)
    end

    private
    def renderers
      @@renderers.map {|r| r.new(@requests, @sheet, @font, @config) }
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
  end
end
