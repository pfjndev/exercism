class BirdCount
  def self.last_week
    @last_week_bird_count = [0, 2, 5, 3, 7, 8, 4]
  end

  def initialize(birds_per_day)
    @last_week_bird_count = birds_per_day
  end

  def yesterday
    @last_week_bird_count[-2]
  end

  def total
    # @last_week_bird_count.sum(0) SIMPLER BUT TAKES LONGER TO RUN (AS OBSERVED)
    @last_week_bird_count.reduce { | sum, birds_per_day | sum + birds_per_day }
  end

  def busy_days
    # Busy days are days when the bird count is equal to or greater than 5.
    # We clean up the array of days we dont care about
    # and then count the number of days remaining.
    @last_week_bird_count.reject { |n_birds| n_birds < 5 } .count
  end

  def day_without_birds?
    @last_week_bird_count.any? { |day_count| day_count == 0}
  end
end
