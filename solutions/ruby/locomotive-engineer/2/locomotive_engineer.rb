class LocomotiveEngineer
  def self.generate_list_of_wagons(*wagons) = wagons

  def self.fix_list_of_wagons(each_wagons_id, missing_wagons)
    first, second, loco, *leftover = each_wagons_id
    [loco, *missing_wagons, *leftover, first, second]
  end

  def self.add_missing_stops(route, **additional_stops)
    {**route, stops: additional_stops.values}
  end

  def self.extend_route_information(route, more_route_information)
    {**route, **more_route_information}
  end
end
