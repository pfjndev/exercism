class Isogram
  
  def self.isogram?(input)
    return true if input.empty?
    
    # Convert to lowercase to handle mixed case
    # Remove spaces and hyphens as they don't count toward duplicates
    # Keep only alphabetic characters for comparison
    # ^a-z means "Any single char in the range a-z"
    chars = input.downcase.gsub(/[^a-z]/, '').chars

    # Check if all characters are unique with array method
    chars == chars.uniq
  end
end
