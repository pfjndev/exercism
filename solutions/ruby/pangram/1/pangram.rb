class Pangram
  
  def self.pangram?(input_string)
    if input_string.empty? 
      return false
    end
    
    letters = input_string
      .gsub(/[^a-zA-Z]/, '')
      .each_char { |c| c.match(/[a-z]/) }
      .upcase
      .chars
      .uniq
      .sort

    #debug letters.count
    #debug "" << letters.join
    return letters.count == 26
  end
end