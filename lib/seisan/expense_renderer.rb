require 'seisan/base_renderer'
require 'date'

module Seisan
  class ExpenseRenderer < BaseRenderer
    def render
      row.call ['立替払サマリー']
      row.call summary_headings
      summary.each do |person, amount|
        row.call [person, amount]
      end
      row.call

      row.call ['立替払明細']
      row.call headings
      lines.each do |line|
        row.call line
      end
    end

    private
    def summary_headings
      %w(氏名 金額)
    end

    def summary
      summary = Hash.new(0)
      requests.each do |entry|
        summary[entry['applicant']] += entry['expense'].inject(0){|r, e| r += e['amount'].to_i }
      end
      summary
    end

    def headings
      %w(日付 立替者 金額 摘要 備考)
    end

    def lines
      lines = []
      requests.each do |entry|
        entry['expense'].each do |expense|
          lines << [expense['date'].to_s, entry['applicant'], expense['amount'], expense['remarks'], expense['notes']]
        end
      end
      lines.sort_by {|line| [Date.parse(line[0]), line[1]] }
    end
  end
end
