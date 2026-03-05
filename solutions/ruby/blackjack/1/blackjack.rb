module Blackjack
    @game_cards = {
      ace: 11,
      two: 2,
      three: 3,
      four: 4,
      five: 5,
      six: 6,
      seven: 7,
      eight: 8,
      nine: 9,
      ten: 10,
      jack: 10,
      queen: 10,
      king: 10
    }

  def self.parse_card(card)
    @game_cards[card.to_sym] || 0
  end

  def self.card_range(card1, card2)
    cards = [card1, card2]
    case cards.reduce(0) { |sum, card| sum + parse_card(card) }
      when 4..11  then "low"
      when 12..16 then "mid"
      when 7..20  then "high"
      when 21     then "blackjack"
    end
  end

  def self.first_turn(card1, card2, dealer_card)
    # Check for pair of aces - always split
    if card1 == "ace" && card2 == "ace"
      return "P"
    end
    
    player_sum = parse_card(card1) + parse_card(card2)
    dealer_value = parse_card(dealer_card)
    
    case player_sum
    when 21  # Blackjack
      # Win automatically unless dealer has ace, face card, or ten
      if ["ace", "ten", "jack", "queen", "king"].include?(dealer_card)
        "S"  # Stand and wait
      else
        "W"  # Automatic win
      end
    when 17..20  # Always stand
      "S"
    when 12..16  # Stand unless dealer has 7 or higher
      if dealer_value >= 7
        "H"  # Hit
      else
        "S"  # Stand
      end
    else  # 11 or lower - always hit
      "H"
    end
  end
end
