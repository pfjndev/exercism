class AssemblyLine
  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour
    success_rate = if @speed >= 1 && @speed <= 4
      1.0
    elsif @speed >= 5 && @speed <= 8
      0.9
    elsif @speed == 9
      0.8
    elsif @speed == 10
      0.77
    else
      0.0
    end
    production = success_rate * (@speed * 221.0)
  end
  
  def working_items_per_minute
    (production_rate_per_hour / 60).to_i
  end
end
