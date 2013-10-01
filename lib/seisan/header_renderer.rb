require 'seisan/base_renderer'

module Seisan
  class HeaderRenderer < BaseRenderer
    def render
      row ["#{organization_name} 精算シート #{target_name}"]
      row ['作成時刻', Time.now.strftime('%Y-%m-%d %X')]
      row
    end

    private
    def target_name
      @config['target']
    end

    def organization_name
      @config['organization'] ? @config['organization']['name'] : ''
    end
  end
end
