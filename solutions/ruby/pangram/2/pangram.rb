class Pangram
  
  def self.pangram?(input_string)
    false if input_string.empty?
    
    letters = input_string
      .gsub(/[^a-zA-Z]/, '')               # 'Delete' any non alphabetic char
      #.each_char { |c| c.match(/[a-z]/) } # Will match '_', which is not ok
      .upcase                              # Avoid case-sensitive
      .chars                               # Array of alphabetic chars
      .uniq                                # Remove duplicates
      #.sort

    #debug letters.count
    #debug "" << letters.join
    return letters.count == 26
  end
end